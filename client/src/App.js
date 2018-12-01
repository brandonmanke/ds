import React, { Component } from 'react';
import './App.css';
import './index.css'

import { Table, Grid, Row, Col, PageHeader } from 'react-bootstrap';
import Feed from './components/Feed'
import Subscriptions from './components/Subscriptions'

var date = new Date()
var time = date.getTime();

var hour = 60*60*1000

  const times = [time, time+(3*hour), time+(6*hour),
                 time+(1*hour), time+(4*hour), time+(7*hour),
                 time+(2*hour), time+(5*hour)]

class App extends Component {

  state = {
    feed: [
      {
        type: 'weather',
        title: "Weather1",
        time: times[0]
      },
      {
        type: 'weather',
        title: "Weather2",
        time: times[1]
      },
      {
        type: 'weather',
        title: "Weather3",
        time: times[2]
      },
      {
        type: 'news',
        title: "News1",
        time: times[3]
      },
      {
        type: 'news',
        title: "News2",
        time: times[4]
      },
      {
        type: 'news',
        title: "News3",
        time: times[5]
      },
      {
        type: 'friends',
        title: "Friend1",
        time: times[6]
      },
      {
        type: 'friends',
        title: "Friend2",
        time: times[7]
      },
    ],
    subscriptions: {
      default: true,
      weather: false,
      news: false,
      friends: false
    }
  }

  defaultSubs = {
    default: true,
    weather: false,
    news: false,
    friends: false
  }

  handleSubscription = (buttons) => {

    console.log(buttons)
    console.log(this.state)

    if(buttons === "default") {
      this.setState(prevState => {
        return {
          default: prevState.subscriptions.default = true,
          weather: prevState.subscriptions.weather = false,
          news: prevState.subscriptions.news = false,
          friends: prevState.subscriptions.friends = false,
        }
      })
    }
    else {
      this.setState(prevState => {
        return {
          default: prevState.subscriptions.default = false,
          buttons: prevState.subscriptions[buttons] = true
        }
      })
    }
  }


  render() {

    console.log(this.state.subscriptions)

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
            <th className="text-center" colSpan="2">
              Feed
            </th>
              {/* <th>Song</th>
              <th>Artist</th>
              <th>Album</th>
              <th>Length</th>
              <th>Votes</th> */}
            </tr>
          </thead>
          {/* <tbody> */}
            <Feed
              feed = {this.state.feed}
              subscriptions = {this.state.subscriptions}
              // Test with {feed just to see} 
            ></Feed>
          {/* </tbody> */}
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
          <Subscriptions
            subscriptions = {this.state.subscriptions}
            changeSubscriptions = {this.handleSubscription}
          />
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
