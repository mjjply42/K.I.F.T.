let spotify_butt = document.querySelector(".spotify-button");

function    startSpotify(order)
{
    console.log(order);
    if (order == "Please open Spotify player\n")
    {
        console.log("dsfghjl");
        let spotify = Application("Spotify");
        spotify.activate();
    }
    else if(order != "Playing\n")
        $(".spotify-button").show();
    $("#response").append('<li>' + order + '<br>' + '</li>');
    var song_synth = window.speechSynthesis;
    var song_speech = new SpeechSynthesisUtterance(order);
    song_speech.lang = 'en-US';
    song_synth.speak(song_speech);
}