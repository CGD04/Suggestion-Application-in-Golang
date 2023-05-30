function login(){
    var data = {
        email : document.getElementById("email").value,
        password:document.getElementById("pw").value
    }
    fetch("login",{
        method:"POST",
        body:JSON.stringify(data),
        headers:{"content-type":"application/json"}
    }).then(response =>{
        if (response.status==200){
            window.location.href = "student.html"
        }else if (response.status==401){
            alert("Enter your Credentials")
        }else{
            throw new Error(response)
        }
    })
}

function Logout(){
    fetch("/logout")
    .then(res => {
        if (res.ok){
            window.open("index.html","_self")
        }else{
            throw new Error(res.statusText)
        }
    }).catch(e =>{
        alert(e)
    })
}