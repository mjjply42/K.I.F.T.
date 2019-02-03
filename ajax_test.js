let load_text = document.getElementById("Load_text");

load_text.addEventListener("click", loadText);

function loadText()
{
    let request = new XMLHttpRequest();
    request.open("GET", "sample.txt", true);
    request.onload = function() 
    {
        if (this.status == 200)
        {
            console.log("Success");
            document.getElementById("field").innerHTML = this.responseText;
        }
    }
    request.send();
}

let del_text = document.getElementById("Del_text");

del_text.addEventListener("click", delText);

function    delText()
{
    document.body.removeChild(document.getElementById('field'));
}