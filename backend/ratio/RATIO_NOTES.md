
Need to refer to https://developer.ratio.me/docs/guides/link-a-new-signing-wallet-to-an-existing-user for flow when users' phone number is already in use.

During wallet + sms auth when *no* user record exists, user mask and user object are returned nil for both APIs.

During wallet + sms auth flow, when user record exists:
* Wallet auth API returns the user mask.
* SMS auth API returns user object.

Q: is there a way to the phone number of an existing user so it doesn't need to be input again for the SMS auth?