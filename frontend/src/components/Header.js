import React from "react";
import {Nav, Navbar} from "react-bootstrap";

function Header() {
    return (
        <Navbar bg="dark" variant="dark">
            <Navbar.Brand href="/">
                <img
                    src={process.env.PUBLIC_URL + "/logo.png"}
                    alt="Logo"
                    height={100}
                />
                <span style={{marginLeft: '10px'}}>GoCalculator</span>
            </Navbar.Brand>
            <Nav className="me-auto">
                <Nav.Link href={"/expressions"}>Expressions</Nav.Link>
                <Nav.Link href="/config">Configuration</Nav.Link>
                <Nav.Link href="/agents">Agents</Nav.Link>
            </Nav>
        </Navbar>
    );
}

export default Header;
