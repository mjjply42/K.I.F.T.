function    checkLights(content)
{
    if (content.substring(0,11).localeCompare(" Turning on") == 0)
    {
        $("#response").append('<li>' + "Turning on" + '<br>' + '</li>');
        $('#image').attr("src", "https://images.cooltext.com/5285731.png");
    }
    else
    {
        $("#response").append('<li>' + "Turning off" + '<br>' + '</li>');
        $('#image').attr("src", "logo1.png");
    }
}