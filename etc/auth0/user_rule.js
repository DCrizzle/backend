// NOTE: this rule is used in conjunction with the @custom
// directives that populate the values in the app_metadata
// and are then just added to the claims on login

function addAttributes(user, context, callback) {
  const claims = {
    "isAuthenticated": true,
    "role": user.app_metadata.role,
    "orgID": user.app_metadata.orgID
  };

  context.idToken["https://folivora.io/jwt/claims"] = claims;
  callback(null, user, context);
}
