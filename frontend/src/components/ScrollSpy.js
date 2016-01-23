import React, { Component, PropTypes } from 'react'
import { connect } from 'react-redux'

export class ScrollSpy extends Component {

  static propTypes = {
    dispatch: PropTypes.func.isRequired
  };

  componentDidMount () {
    const { dispatch } = this.props
    setTimeout(() => {
      window.scrollTo(0, 10000)
    }, 0)
    window.onscroll = (event) => {
      const doc = document.documentElement
      const top = (window.pageYOffset || doc.scrollTop)  - (doc.clientTop || 0)
      console.log(top)
      dispatch(event)
    }
  }

  render () {
    return <div />
  }

}

export default connect()(ScrollSpy)
