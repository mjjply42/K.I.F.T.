let spotify_butt = document.querySelector(".spotify-button");

function    startSpotify()
{
    $(".spotify-button").show();
    $("#response").append('<li>' + "Please Sign In First" + '</li>');
    var song_synth = window.speechSynthesis;
    var song_speech = new SpeechSynthesisUtterance("Sign In");
    song_speech.lang = 'en-US';
    song_synth.speak(song_speech);
}