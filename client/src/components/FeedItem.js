import React from 'react'

class FeedItem extends React.Component {

    render() {

        const {
            title,
            time,
        } = this.props

        // var date = new Date(time)
        // var hour = date.getHours();
        // var minutes = date.getMinutes();
        // var ampm = "AM";

        // if(hour > 12) {
        //     hour = hour - 12;
        //     ampm = "PM";
        // } else if(hour === 0) {
        //     hour = 12;
        // }

        // const realTime = hour + ":" + minutes + ampm;

        return(
            <tr>
                <td>{title}</td>
                <td>{time}</td>
            </tr>
        )
    }
}

export default FeedItem