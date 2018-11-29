import React from 'react';

const weather = [
    {
      title: "Weather1",
      time: "9:00 PM"
    },
    {
        title: "Weather2",
        time: "10:00 PM"
    },
    {
        title: "Weather3",
        time: "11:00 PM"
    },
  ]

class Feed extends React.Component {

    
      
    sortFeed(entryA, entryB) {
        //TODO
    }

    render() {

        const list = weather;

        return(
            <p>
            {list.map((weather, index) => (
                <div>{weather.title}
                {weather.time}</div>
            ))}
            </p>
        );
    }
}

export default Feed;