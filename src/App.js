import React from 'react';
import Container from 'react-bootstrap/Container';
import Card from 'react-bootstrap/Card';
import Col from 'react-bootstrap/Col';
import Row from 'react-bootstrap/Row';
import Navbar from 'react-bootstrap/Navbar';
import Stack from 'react-bootstrap/Stack';
import './App.css';

function App() {
  return (
    <div className="App">
      <Navbar className="bg-body-tertiary">
        <Container>
          <Navbar.Brand href="#home">
            <img
              src="/img/octopus.svg"
              width="50"
              height="50"
              className="d-inline-block align-top"
              alt="React Bootstrap logo"
            />
          </Navbar.Brand>
        </Container>
      </Navbar>
      <header className="App-header">
        <Container>
          <Row xs={1} md={2}>
            <Col>
              <div className="board open">
                <h2 className='title-text'>Open <span className='count'>2</span></h2>
                <Row className="g-4 flex-column card_wrap">
                  {Array.from({ length: 2 }).map((_, idx) => (
                    <Col className='p-0' key={idx}>
                      <Card className='open text-start'>
                        {/* <Card.Img variant="top" src="holder.js/100px160" /> */}
                        <Card.Body>
                          <Card.Title>Research new security protocols</Card.Title>
                          <Card.Text>
                            This task is in here because it has been identified as necessary.
                          </Card.Text>
                          <Stack direction="horizontal" gap={2}><div className='tag info'>Devops</div></Stack>
                        </Card.Body>
                      </Card>
                    </Col>
                  ))}
                </Row>
              </div>
            </Col>
            <Col>
              <div className="board close">
                <h2 className='title-text'>Closed <span className='count'>2</span></h2>
                <Row className="g-4 flex-column card_wrap">
                  {Array.from({ length: 2 }).map((_, idx) => (
                    <Col className='p-0' key={idx}>
                      <Card className='close text-start'>
                        {/* <Card.Img variant="top" src="holder.js/100px160" /> */}
                        <Card.Body>
                          <Card.Title>Finalize UI design</Card.Title>
                          <Card.Text>
                            The Review column because the design has been completed and is awaiting approval from the UX team
                          </Card.Text>
                          <Stack direction="horizontal" gap={2}><div className='tag status'>UX design</div><div className='tag status'>UI design</div></Stack>
                        </Card.Body>
                      </Card>
                    </Col>
                    ))}
                  </Row>
              </div>
            </Col>
          </Row>
        </Container>
      </header>
    </div>
  );
}

export default App;
