export const SAGA_INIT_APP = 'SAGA_INIT_APP'

export function createAction(type, payload) {
    const action = { type, payload }
    if (payload instanceof Error) {
        action.error = true
    }
    return action
}

const actions = {}

actions.initApp = () => {
  return createAction(SAGA_INIT_APP, {})
}

export default actions