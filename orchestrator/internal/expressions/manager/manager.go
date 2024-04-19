package manager

import (
	"github.com/chessnok/GoCalculator/orchestrator/internal/db"
	"github.com/chessnok/GoCalculator/orchestrator/pkg/rabbit/queue"
	"github.com/chessnok/GoCalculator/orchestrator/pkg/result"
	pb "github.com/chessnok/GoCalculator/proto"
	"github.com/rabbitmq/amqp091-go"
	"time"
)

type TasksManager struct {
	db         *db.Postgres
	producer   *queue.Producer
	consumer   *queue.Consumer
	calcConfig *pb.Config
	stop       chan struct{}
}

func (tm *TasksManager) SendTasksToQueue() {
	tasks, err := tm.db.Tasks.SelectTasksToSendToQueue(tm.calcConfig)
	if err != nil {
		return
	}
	for _, task := range tasks {
		err := tm.producer.SendJson(task)
		if err == nil {
			tm.db.Tasks.UpdateTaskStatus(task.Id, "in_queue")
		}
	}
}

func (tm *TasksManager) OnNewResult(delivery *amqp091.Delivery) {
	res, err := result.ResultFromDelivery(delivery)
	if err != nil {
		return
	}
	tm.db.Tasks.TaskResult(res.Id, res.Result, res.IsErr)
}
func NewTasksManager(db *db.Postgres, producer *queue.Producer, consumer *queue.Consumer, calcConfig *pb.Config) *TasksManager {
	return &TasksManager{
		db:         db,
		producer:   producer,
		consumer:   consumer,
		calcConfig: calcConfig,
	}
}
func (tm *TasksManager) Start() {
	go tm.consumer.Consume(tm.OnNewResult)
	go func() {
		for {
			select {
			case <-tm.stop:
				tm.consumer.Stop <- struct{}{}
				return
			default:
				tm.SendTasksToQueue()
				time.Sleep(time.Second / 4)
			}
		}
	}()
}
