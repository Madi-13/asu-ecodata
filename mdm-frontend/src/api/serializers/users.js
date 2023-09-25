export const deserializeUser = (payload) => {
  return {
    id: payload.id,
    firstName: payload.firstname,
    lastName: payload.lastname,
    username: payload.username,
    email: payload.email,
    emailVerified: true,
    enabled: true,
    credentials: [
      {
        type: 'password',
        value: payload.password,
        temporary: false
      }
    ],
    clientRoles: {
      suppression: ['admin']
    },
    groups: [],
    attributes: {}
  }
}

export const serializeUser = (payload) => ({
  id: payload.id,
  firstname: payload.firstName,
  email: payload.email,
  roles: [],
  lastname: payload.lastName,
  username: payload.username,
  password: payload?.password || null,
  password_rec: null
})
