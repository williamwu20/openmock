import { createStore, applyMiddleware, compose } from 'redux'
import { routerMiddleware } from 'react-router-redux'
import createSagaMiddleware from 'redux-saga'

import { createBrowserHistory } from 'history'
import rootReducer from './reducers'
import { defineSagaListeners } from './sagas'

const sagaMiddleware = createSagaMiddleware()

export const history = createBrowserHistory()

const middleware = [routerMiddleware(history), sagaMiddleware]
const enhancers = []

const composedEnhancers = compose(applyMiddleware(...middleware), ...enhancers)

export default (initialState) => {
  const store = createStore(rootReducer, initialState, composedEnhancers)
  sagaMiddleware.run(defineSagaListeners)
  return store
}
