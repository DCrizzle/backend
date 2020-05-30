function (user, context, callback) {
  const axios = require('axios@0.19.2');

  const auth0_options = {
		method: 'POST',
    url: `https://${auth0.domain}/oauth/token`,
    headers: {
			'Content-Type': 'application/json'
		},
    data: `
			{
				'client_id':'${configuration.RULE_APP_CLIENT_ID}',
				'client_secret':'${configuration.RULE_APP_CLIENT_SECRET}',
				'audience':'https://tbdbackend.io/',
				'grant_type':'client_credentials'
			}`
	};
	console.log('auth0_options:', auth0_options);

  axios(auth0_options)
  .then( auth0_response => {
		console.log('auth0_response:', auth0_response);
    const access_token = auth0_response.data.access_token;

		const dgraph_options = {
			method: 'POST',
			url: configuration.TBD_BACKEND_URL,
			headers: {
				'Content-Type': 'application/json',
				'X-Auth0-Token': access_token
			},
			data: {
				query: `
					query GetUser($email: String!) {
	          getUser(email: $email) {
	            firstName
							lastName
							role
							orgID
	          }
	        }`,
				variables: {
					email: user.email
				}
			}
		};
		console.log('dgraph_options:', dgraph_options);

		axios(dgraph_options)
		.then( dgraph_response => {
			console.log('dgraph_response:', dgraph_response);

			context.idToken['https://tbd.io/jwt/claims'] = {
				'role': dgraph_response.data.role,
				'orgID': dgraph_response.data.orgID
			};
			console.log('context:', context);
		})
		.catch( dgraph_error => {
			console.log('dgraph_error:', dgraph_error);
		});
  })
  .catch( auth0_error => {
    console.log('auth0_error', auth0_error);
  });

  return callback(null, user, context);
}

// notes:
// https://www.apollographql.com/blog/4-simple-ways-to-call-a-graphql-api-a6807bcdb355
