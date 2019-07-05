function stateCheck(jsFilePath) {
    let split_arr = jsFilePath.split(".");
    let name = split_arr[0];
    let script = document.querySelector("."+name);
    if(script.id == "active")
        script.id = "inactive";
    else
        script.id = "active"
}