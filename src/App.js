import React from 'react';
import Container from 'react-bootstrap/Container';
import Card from 'react-bootstrap/Card';
import Col from 'react-bootstrap/Col';
import Row from 'react-bootstrap/Row';
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <Container>
          <Row xs={1} md={2}>
            <Col>
              <div className="board open">
                <h2 className='title-text'>Open</h2>
                <Row className="g-4 flex-column card_wrap">
                  {Array.from({ length: 2 }).map((_, idx) => (
                    <Col key={idx}>
                      <Card className='open'>
                        {/* <Card.Img variant="top" src="holder.js/100px160" /> */}
                        <Card.Body>
                          <Card.Title>Research new security protocols</Card.Title>
                          <Card.Text>
                            This task is in here because it has been identified as necessary.
                          </Card.Text>
                        </Card.Body>
                      </Card>
                    </Col>
                  ))}
                </Row>
              </div>
            </Col>
            <Col>
              <div className="board close">
                <h2 className='title-text'>Closed</h2>
                <Row className="g-4 flex-column card_wrap">
                  {Array.from({ length: 2 }).map((_, idx) => (
                    <Col key={idx}>
                      <Card className='close'>
                        {/* <Card.Img variant="top" src="holder.js/100px160" /> */}
                        <Card.Body>
                          <Card.Title>Finalize UI design</Card.Title>
                          <Card.Text>
                            The Review column because the design has been completed and is awaiting approval from the UX team
                          </Card.Text>
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
