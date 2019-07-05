function stateCheck(jsFilePath) {
    let split_arr = jsFilePath.split(".");
    let name = split_arr[0];
    let script = document.querySelector("."+name);
    if(script.id == "active")
        return 0;
    else
    {
        script.id = "active";
        return 1;
    }
}