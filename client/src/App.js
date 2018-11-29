import React, { Component } from 'react';
import './App.css';
import './index.css'

import { Table, Grid, Row, Col, PageHeader, Button } from 'react-bootstrap';
import Feed from './components/Feed'

class App extends Component {
  render() {
    return (

      <Grid>
      <PageHeader>
        <Row>
          <Col md={3}>PUB -- SUB</Col>
          <Col md={3}>
            <small>Server by Brandon Manke</small>
            {/* <small> Room Code: {this.state.roomCode}</small> */}
            {/* <small> Room Code: {roomCode}</small> */}
          </Col>
          <Col md={3}>
            {/* <small>{this.props.room.roomName}</small> */}
          </Col>
          <Col md={3}>
          <small>UI and Sockets by Rob Putyra</small>
            {/* <Button onClick={this.debugAdd}>Add song</Button> */}
          </Col>
        </Row>
      </PageHeader>
      <Row>
        {/* <SongSearch /> */}
      </Row>
      <Row>
        <Col md={8}>
        <Table striped bordered condensed hover responsive>
          <thead>
            <tr>
            <th className="text-center">
              Feed
            </th>
              {/* <th>Song</th>
              <th>Artist</th>
              <th>Album</th>
              <th>Length</th>
              <th>Votes</th> */}
            </tr>
          </thead>
          <tbody>
            <Feed></Feed>
          </tbody>
          {/* <SongList songs={songs} vote={this.vote} /> */}
          {/* <tbody>
            {songs.map((song, index) => (
              // console.log(`Song: ${song}`)
              <Song
                title={song.data.name}
                artist={song.data.artist.name}
                album={song.data.album.name}
                songLength={'0'}
                votes={song.data.votes}
                id={song.data.id}
                key={song.key}
                songKey={song.key}
                index={index}
                vote={this.vote}
              />
            ))}
          </tbody> */}
        </Table>
        </Col>
        {/* TODO: These buttons should be in their own component class */}
        <Col className="text-center" md={4}>
          <Row className="with-margin">
            <Button block>Default</Button>
          </Row>
          <Row className="with-margin">
            <Button block bsStyle="primary">Primary</Button>
          </Row>
          <Row className="with-margin">
            {/* Indicates a successful or positive action */}
            <Button block bsStyle="success">Success</Button>
          </Row>
          <Row className="with-margin">
            {/* Contextual button for informational alert messages */}
            <Button block bsStyle="info">Info</Button>
          </Row>
          <Row className="with-margin">
            {/* Indicates caution should be taken with this action */}
            <Button block bsStyle="warning">Warning</Button>
          </Row>
          <Row className="with-margin">
            {/* Indicates a dangerous or potentially negative action */}
            <Button block bsStyle="danger">Danger</Button>
          </Row>
        </Col>
      </Row>
    </Grid>

      // <div className="App">
      //   <header className="App-header">
      //     <img src={logo} className="App-logo" alt="logo" />
      //     <p>
      //       Edit <code>src/App.js</code> and save to reload.
      //     </p>
      //     <a
      //       className="App-link"
      //       href="https://reactjs.org"
      //       target="_blank"
      //       rel="noopener noreferrer"
      //     >
      //       Learn React
      //     </a>
      //   </header>
      // </div>
    );
  }
}

export default App;
