import  React  from 'react'
import {Launcher} from 'react-chat-window'
import axios from 'axios'


class Messenger extends React.Component {
 
    constructor() {
      super();
      this.state = {
        messageList: [],
        usersOnline:[]
      };
    }

    componentDidUpdate() {
      axios.get("http://127.0.0.1:8000/onlineusers")
        .then(response => {
          return this.setState({ usersOnline: response });
        })
    }
   
    _onMessageWasSent(message) {
      this.setState({
        messageList: [...this.state.messageList, message]
      })
    }
   
    _sendMessage(text) {
      if (text.length > 0) {
        this.setState({
          messageList: [...this.state.messageList, {
            author: 'them',
            type: 'text',
            data: { text }
          }]
        })
      }
    }
   
    render() {
      return (<div>
        <Launcher
          agentProfile={{
            teamName: this.state.usersOnline,
          }}
          onMessageWasSent={this._onMessageWasSent.bind(this)}
          messageList={this.state.messageList}
          showEmoji
        />
      </div>)
    }
  }

  export default Messenger