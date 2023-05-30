function signUp(){
    // data to be sent to the POST request
    var data = {
        fname : document.getElementById("fname").value,
        lname : document.getElementById("lname").value,
        email : document.getElementById("email").value,
        password : document.getElementById("pw1").value,
        password2 : document.getElementById("pw2").value,
    }
    
    if (data.password === data.password2){
        fetch("/signup",{
            method:"POST",
            body:JSON.stringify(data),
            headers:{"Content-type":"application.json; charset=UTF-8"}
        }).then(response => {
            if (response.status == 201){
                window.open("index.html","_self")
            }
        })
    }

}