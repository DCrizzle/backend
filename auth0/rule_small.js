// NOTE: this rule is used in conjunction with the @custom
// directives that populate the values in the app_metadata
// and are then just added to the claims on login

function (user, context, callback) {
  const namespace = "https://folivora.io/jwt/claims";
  context.idToken[namespace] =
    {
			isAuthenticated: true,
    	role: user.app_metadata.role,
    	orgID: user.app_metadata.orgID,
    };

  return callback(null, user, context);
}
