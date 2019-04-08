/**
 * This is an example of a basic node.js script that performs
 * the Authorization Code oAuth2 flow to authenticate against
 * the Spotify Accounts.
 *
 * For more information, read
 * https://developer.spotify.com/web-api/authorization-guide/#authorization_code_flow
 */

var express = require('express'); // Express web server framework
var request = require('request'); // "Request" library
var cors = require('cors');
var querystring = require('querystring');
var cookieParser = require('cookie-parser');

var client_id = 'f137e890f3734bd38fdcf1d980158139'; // Your client id
var client_secret = '119bb13ee90043919bb53c48baf17fa9'; // Your secret
var redirect_uri = 'http://localhost:8888/callback/'; // Your redirect uri

/**
 * Generates a random string containing numbers and letters
 * @param  {number} length The length of the string
 * @return {string} The generated string
 */
var generateRandomString = function(length) {
  var text = '';
  var possible = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';

  for (var i = 0; i < length; i++) {
    text += possible.charAt(Math.floor(Math.random() * possible.length));
  }
  return text;
};

var app = express();

app.use(express.static(__dirname + '/public'))
   .use(cors())
   .use(cookieParser());

app.get('/login', function(req, res) {
  // your application requests authorization
  var scope = 'user-read-private user-read-email user-read-playback-state';
  var url = 'https://accounts.spotify.com/authorize?' +
  querystring.stringify({
    response_type: 'code',
    client_id: client_id,
    scope: scope,
    redirect_uri: redirect_uri,
  })
  console.log(url);
  res.redirect(url);
});

app.get('/callback', function(req, res) {

  // your application requests refresh and access tokens
  // after checking the state parameter
  console.log(req.query)
  var code = req.query.code || null;

  var authOptions = {
    url: 'https://accounts.spotify.com/api/token',
    form: {
      code: code,
      redirect_uri: redirect_uri,
      grant_type: 'authorization_code'
    },
    headers: {
      'Authorization': 'Basic ' + (new Buffer(client_id + ':' + client_secret).toString('base64'))
    },
    json: true
  };
  //After I grab back the access_token I can now make a get request with it.
  request.post(authOptions, function(error, response, body) {
      console.log("this is body: ", body);
      // console.log("this is access Token: ", body.access_token)
  });
});


console.log('Listening on 8888');
app.listen(8888);
