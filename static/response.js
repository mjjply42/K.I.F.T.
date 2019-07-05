function    sendResponse(command, content)
{
    let speak = 0;
    let email = "test@gmail.com";
    let duty = command;

    if(command == "Email")
    {
        $(".client-input").show();
        $("#response").append('<li>' + content + '</li>');
        var voice_synth = window.speechSynthesis;
        var voice_speech = new SpeechSynthesisUtterance("Please type email address");
        voice_speech.lang = 'en-US';
        voice_synth.speak(voice_speech);
        $(".client-input").on('keyup', function (event) {
        if (event.keyCode == 13) {
        email = $(".client-input").val();
        $(".client-input").val("");
        $(".client-input").hide();
            $.ajax({
                url: "/response",
                method: "POST",
                data: {email, speak, duty},
                success: function (data) {
                console.log(data)
            }});
        }
        })
    }
    else if(command == "Weather")
    {
        speak = 1;
        $("#response").append('<li>' + content + '</li>');
        var voice_synth = window.speechSynthesis;
        var voice_speech = new SpeechSynthesisUtterance("Tell me what city");
        voice_speech.lang = 'en-US';
        voice_synth.speak(voice_speech);
            $.ajax({
                url: "/response",
                method: "POST",
                data: {speak,duty},
                success: function (data) {
                console.log(data)
            }});

    }
    else if(command == "Event")
    {
        speak = 1;
        $("#response").append('<li>' + content + '</li>');
        var voice_synth = window.speechSynthesis;
        var voice_speech = new SpeechSynthesisUtterance("Tell me what city");
        voice_speech.lang = 'en-US';
        voice_synth.speak(voice_speech);
            $.ajax({
                url: "/response",
                method: "POST",
                data: {speak,duty},
                success: function (data) {
                console.log(data)
            }});

    }
    else if(command == "Define")
    {
        speak = 1;
        $("#response").append('<li>' + content + '</li>');
        var voice_synth = window.speechSynthesis;
        var voice_speech = new SpeechSynthesisUtterance("Tell me the word");
        voice_speech.lang = 'en-US';
        voice_synth.speak(voice_speech);
            $.ajax({
                url: "/response",
                method: "POST",
                data: {speak,duty},
                success: function (data) {
                console.log(data)
            }});

    }
}