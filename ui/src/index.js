import React from 'react';
import ReactDOM from 'react-dom';
import { Provider, connect } from "react-redux";
import store, { history } from "./store";
import AppBar from './components/AppBar';

ReactDOM.render(
  <React.StrictMode>
      <Provider store={store({})}>
        <AppBar />
      </Provider>
  </React.StrictMode>,
  document.getElementById('root')
);
