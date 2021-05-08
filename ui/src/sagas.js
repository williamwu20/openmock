import { takeLatest } from 'redux-saga/effects'
import {
  SAGA_INIT_APP
} from './actions'

export function* handleInit({}) {
    try {
    } catch (error) {
      console.log(error)
    }
  }

export function* defineSagaListeners() {
    yield [
      takeLatest(SAGA_INIT_APP, handleInit)
    ]
  }

