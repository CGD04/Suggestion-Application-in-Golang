
window.onload = function(){
    fetch('/courses')
    .then(response => response.text())
    .then(data => showCourses(data))

    
}

function showCourses(data) {
    courses =JSON.parse(data)
    courses.forEach(course => {
        newRow(course)
    });
}
function addCourse() {
    _data = getFormData()

    if (_data.id == "" || _data.courseName == ""){
        alert("Course Id cannot be empty")
        return 
    }
    fetch("/course",{
        method:"POST",
        body:JSON.stringify(_data),
        header:{"Content-type":"application/json; charset=UTF-8"}
    }).then(response =>{
        if (response.ok){
            fetch("/course/"+_data.id)
            .then(response1 => response1.text())
            .then(info => showCourse(info))
        }else if(response.status == 409){
            errMessage()
        }
    })
}

function errMessage() {
    // parentDiv = document.getElementById("container")
    // newDiv = document.createElement("div")
    // parentDiv.appendChild(newDiv)
    // para = document.createElement("p")
    // newDiv.appendChild(para)
    // para.innerHTML = "The student is already registered"
    // para.setAttribute("style","color:red;")
    // setTimeout(function(){
    //     para.innerHTML = "";
    // },700);
    alert("The student is already registered")
}
function showCourse(info){
    data = JSON.parse(info)
    newRow(data)
}

function newRow(data) {
    var table = document.getElementById("myTable")
    var row = table.insertRow(table.length)
    var td =  []

    for (i = 0; i < table.rows[0].cells.length;i++){
        td[i] = row.insertCell(i)
    }
    td[0].innerHTML = data.id
    td[1].innerHTML = data.courseName
    td[2].innerHTML = '<input type="button" onclick="deleteCourse(this)" value="Delete" id="button-1"/>'
    td[3].innerHTML = '<input type="button" onclick="updateCourse(this)" value="Update" id="button-2"/>'
    resetForm()
}

function deleteCourse(r) {
    selectedRow = r.parentElement.parentElement
    id = selectedRow.cells[0].innerHTML
    if (confirm("Are you sure you want to delete the course with id "+id)){
        fetch("/course/"+id,{
            method:"DELETE",
            headers:{"Content-type":"application/json; charset=UTF-8"}
        }).then(response =>{
            if (response.ok){
                var rowIndex = selectedRow.rowIndex
                if (rowIndex > 0){
                    selectedRow.parentElement.deleteRow(selectedRow.rowIndex)
                }
                selectedRow = null
            }else{
                throw new Error(response.statusText)
            }
        }).catch(e =>{
            alert(e)
        })
    }
}
function getFormData() {
    var data = {
        id: document.getElementById("cid").value,
        courseName: document.getElementById("cname").value
    }
    return data
}
function resetForm() {
    document.getElementById("cid").value = "";
    document.getElementById("cname").value = "";
}
function updateCourse(r) {
    selectedRow = r.parentElement.parentElement
    document.getElementById("cid").value = selectedRow.cells[0].innerHTML;
    document.getElementById("cname").value = selectedRow.cells[1].innerHTML;
    data = getFormData()
    button = document.getElementById("button-add")
    cid = selectedRow.cells[0].innerHTML
    button.innerHTML = "Update"
    button.setAttribute("onclick", "update(cid)")
}

function update(id) {
    var data = getFormData()
    fetch("/course/"+id, {
        method:"PUT",
        body:JSON.stringify(data),
        headers:{"Content-type":"application/json; charset=UTF-8"}
    }).then(response => {
        if (response.ok){
            selectedRow.cells[0].innerHTML = data.id;
            selectedRow.cells[1].innerHTML = data.courseName;
            button = document.getElementById("button-add");
            button.innerHTML = "Add";
            button.setAttribute("onclick","addCourse()");
            selectedRow = null;
            document.getElementById("cid").innerHTML = "";
            resetForm()
        }    
    }).catch(e =>{
        alert(e)
    })
}


