function    sendResponse(command, content)
{
    let speak = 0;
    let email = "test@gmail.com";
    let duty = command;
    let value = 'a';

    if(command == "email")
    {
        $(".client-input").show();
        $(".client-input").focus();
        $("#response").append('<li>' + content + '<br>' + '</li>');
        var voice_synth = window.speechSynthesis;
        var voice_speech = new SpeechSynthesisUtterance("Please type email address");
        voice_speech.lang = 'en-US';
        voice_synth.speak(voice_speech);
        $(".client-input").on('keyup', function (event) {
        if (event.key == "Enter") {
        email = $(".client-input").val();
        $(".client-input").val("");
        $(".client-input").hide();
        $(".client-input").unbind();
            $.ajax({
                url: "/response",
                method: "POST",
                data: {email, speak, duty},
                success: function (data) {
                    $("#response").append('<li>' + "Email Sent!" + '<br>' +'</li>');
                    var voice_synth = window.speechSynthesis;
                    var voice_speech = new SpeechSynthesisUtterance("Email Sent");
                    voice_speech.lang = 'en-US';
                    voice_synth.speak(voice_speech);
            }});
        }
        })
        return ;
    }
    else if(command === "weather")
    {
        $(".client-input").show();
        $(".client-input").focus();
        $("#response").append('<li>' + content + '<br>' + '</li>');
        var weather_synth = window.speechSynthesis;
        var weather_speech = new SpeechSynthesisUtterance("Please input the city");
        weather_speech.lang = 'en-US';
        weather_synth.speak(weather_speech);
        $(".client-input").on('keyup', function () {
            if (event.key == "Enter") {
            value = $(".client-input").val();
            $(".client-input").val("");
            $(".client-input").hide();
            weather_synth.cancel();
            $(".client-input").unbind();
            if(value != "")
            {
                $.ajax({
                    url: "/response",
                    method: "GET",
                    cache: false,
                    data: {value, duty},
                    success: function (data) {
                    $("#response").append('<li>' + data + '<br>' + '</li>');
                    return ;
                }});     
            }           
            }
        })
    }
    else if(command == "event")
    {
        $(".client-input").show();
        $(".client-input").focus();
        $("#response").append('<li>' + content + '<br>' + '</li>');
        var event_synth = window.speechSynthesis;
        var event_speech = new SpeechSynthesisUtterance("Please input the city");
        event_speech.lang = 'en-US';
        event_synth.speak(event_speech);
        $(".client-input").on('keyup', function (event) {
            if (event.key == "Enter") {
                value = $(".client-input").val();
            $(".client-input").val("");
            $(".client-input").hide();
            event_synth.cancel();
            $(".client-input").unbind();
            if(value != "")
            {
                $.ajax({
                    url: "/response",
                    method: "GET",
                    cache: false,
                    data: {value, duty},
                    success: function (data) {
                    $("#response").append('<li>' + data + '<br>' + '</li>');
                }});
            }
        }
            })
            return;
    }
    else if(command == "define")
    {
        speak = 1;
        $("#response").append('<li>' + content + '<br>' + '</li>');
        var define_synth = window.speechSynthesis;
        var define_speech = new SpeechSynthesisUtterance("Type the word");
        define_speech.lang = 'en-US';
        define_synth.speak(define_speech);
        $(".client-input").on('keyup', function (event) {
            if (event.key == "Enter") {
                value = $(".client-input").val();
            $(".client-input").val("");
            $(".client-input").hide();
            $(".client-input").unbind();
                $.ajax({
                    url: "/response",
                    method: "POST",
                    data: {value, duty},
                    success: function (data) {
                    $("#response").append('<li>' + data + '<br>' + '</li>');
                }});
            }
            })
            return ;
    }
}