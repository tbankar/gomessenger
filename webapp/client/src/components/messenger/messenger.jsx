import  React  from 'react'
import {Launcher} from 'react-chat-window'

class Messenger extends React.Component {
 
    constructor() {
      super();
      this.state = {
        messageList: []
      };
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
            teamName: 'react-chat-window',
          }}
          onMessageWasSent={this._onMessageWasSent.bind(this)}
          messageList={this.state.messageList}
          showEmoji
        />
      </div>)
    }
  }

  export default Messenger