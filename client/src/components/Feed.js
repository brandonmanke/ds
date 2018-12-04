import React from 'react';
import FeedItem from './FeedItem';


class Feed extends React.Component {

    
      
    sortFeed(entryA, entryB) {
        //TODO
    }

    render() {

        const {
            feed
        } = this.props

        if(feed==='') {
            console.log("feed is null")
            return <tbody></tbody>
        }
         else {
             console.log(feed)
            feed.sort((a,b) => {
                return a.time - b.time;
            })

            return(
                <tbody>
                {feed.map((item, index) => (
                    <FeedItem
                        title = {item.title}
                        time = {item.time}
                        key = {index}
                    ></FeedItem>
                ))}
                </tbody>
            );


         }
        

        
    }
}

export default Feed;