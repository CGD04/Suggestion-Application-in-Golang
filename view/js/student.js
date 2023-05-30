
window.onload = function() {
    fetch('/students') 
    .then(response => response.text())
    .then(data => showStudents(data))
}

function showStudents(a) {
    const students = JSON.parse(a)
    students.forEach(stud => {
      newRow(stud)
    })
}
function addStudent(){
    var data = {
        StdID : parseInt(document.getElementById("sid").value),
        FirstName : document.getElementById("fname").value,
        LastName: document.getElementById("lname").value,
        Email:document.getElementById("email").value,
    }
    //second argument is not needed if get method is used
    var sid = data.StdID
    if (isNaN(sid)) {
        alert("Enter valid student ID")
        return
    }else if (data.Email == ""){
        alert("Email cannot be empty")
        return
    }else if (data.FirstName == ""){
        alert("first name cannot be empty")
        return
    }
    fetch("/student", {
        method: "POST",
        body:JSON.stringify(data),
        headers: {"Content-type":"application/json; charset=UTF-8"}
    }).then(response => {
        if (response.ok) {
            fetch('/student/'+sid)
            .then(response => response.text())
            .then(data => showStudent(data))
            resetForm()
        }else{
            throw new Error(response.statusText)
        }
    }).catch(e => {
        alert(e)
    })
}

function showStudent(a) {
    const student = JSON.parse(a)
    newRow(student)
}

function resetForm() {
    document.getElementById("sid").value = "";
    document.getElementById("fname").value = "";
    document.getElementById("lname").value = "";
    document.getElementById("email").value = "";

}

function newRow(student) {

    //find a table element with id='myTable':
    var table = document.getElementById("myTable");

    var row = table.insertRow(table.length);

    var td = []
    for (i = 0; i < table.rows[0].cells.length; i++){
        td[i] = row.insertCell(i)
    }
    td[0].innerHTML = student.StdID
    td[1].innerHTML = student.FirstName
    td[2].innerHTML = student.LastName
    td[3].innerHTML = student.Email
    td[4].innerHTML = '<input type="button" onclick="deleteStudent(this)" value="delete" id="button-1"/>'
    td[5].innerHTML = '<input type="button" onclick="updateStudent(this)" value="Update" id="button-2"/>'

}

function updateStudent (r) {
    selectedRow = r.parentElement.parentElement;
    document.getElementById("sid").value = selectedRow.cells[0].innerHTML;
    document.getElementById("fname").value = selectedRow.cells[1].innerHTML;
    document.getElementById("lname").value = selectedRow.cells[2].innerHTML;
    document.getElementById("email").value = selectedRow.cells[3].innerHTML;

    addButton = document.getElementById("button-add")
    sid = selectedRow.cells[0].innerHTML
    addButton.innerHTML = "Update"
    addButton.setAttribute("onclick", "update(sid)");
}


function getFormData() {
    var data = {
        stdid : parseInt(document.getElementById("sid").value),
        firstname : document.getElementById("fname").value,
        lastname: document.getElementById("lname").value,
        email:document.getElementById("email").value,
    }
    return data
}

function update(sid) {
    var data = getFormData()
    fetch("/student/"+sid,{
        method: "PUT",
        body:JSON.stringify(data),
        headers: {"Content-type":"application/json; charset=UTF-8"}
    }
    ).then(response => {
        if (response.ok) {
            selectedRow.cells[0].innerHTML = data.stdid;
            selectedRow.cells[1].innerHTML = data.firstname;
            selectedRow.cells[2].innerHTML = data.lastname;
            selectedRow.cells[3].innerHTML = data.email;
            addButton = document.getElementById("button-add")
            addButton.innerHTML = "Add"
            addButton.setAttribute("onclick", "AddStudent()");
            selectedRow = null;
            resetForm()
            
        }else{
            throw new Error(response.statusText)
        }
    }).catch(e => {
        alert(e)
    })
}


var selectedRow
function deleteStudent(r) {
    selectedRow = r.parentElement.parentElement
    sid = selectedRow.cells[0].innerHTML
    if (confirm("Are you sure you want to delete this student?")){
        fetch("/student/"+sid,{
            method: "DELETE",
            headers: {"Content-type":"application/json; charset=UTF-8"}
        }
        ).then(response => {
            if (response.ok) {
                var rowIndex = selectedRow.rowIndex;
                if (rowIndex > 0){
                    document.getElementById("myTable").deleteRow(rowIndex);
                }
                selectedRow = null
            }else{
                throw new Error(response.statusText)
            }
        }).catch(e => {
            alert(e)
        })
    }
}
