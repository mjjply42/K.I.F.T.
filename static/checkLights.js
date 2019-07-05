function    checkLights(content)
{
if (content.substring(0,11).localeCompare(" Turning on") == 0) 
    $('#image').attr("src", "https://images.cooltext.com/5285731.png");
else
      $('#image').attr("src", "logo1.png");
}