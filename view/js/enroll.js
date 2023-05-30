window.onload = function(){
    fetch("/students")
    .then(res => res.text())
    .then(data => getStudents(data))
    .catch(e =>{
        alert(e)
    }),
    fetch("/courses")
    .then(resp => resp.text())
    .then(data => showCourses(data))
    .catch(e => {
        alert(e)
    }),
    fetch("/showEnrolledAll")
    .then(response =>response.text())
    .then(data => showEnrolled(data))
}

function addEnroll() {
    var data = {
        stdid: parseInt(document.getElementById("sid").value),
        courseid: document.getElementById("cid").value
    }
    fetch("/enroll",{
        method:"POST",
        body:JSON.stringify(data),
        headers:{"content-type":"application/json; charset=utf-8"}
    }).then(res => {
        if (res.status == 201) {
            fetch("/showEnrolled/"+data.stdid +"/"+data.courseid,)
            .then(resp => resp.text())
            .then(data => shwoDatum(data))
        }else{
            throw new Error(res.statusText)
        }
    }).catch(e =>{
        alert(e)
    })
}


function shwoDatum(data) {
    let info = JSON.parse(data)
    newRow(info)
}

function showEnrolled(data) {
    let enrolled = JSON.parse(data)
    enrolled.forEach(enrol => {
        newRow(enrol)
    });
}
function newRow(info) {
    var table = document.getElementById("myTable")
    var row = table.insertRow(table.length)
    var td =  []
    for (i = 0; i < table.rows[0].cells.length;i++){
        td[i] = row.insertCell(i)
    }
    td[0].innerHTML = info.stdid
    td[1].innerHTML = info.courseid
    td[2].innerHTML = info.date
    td[3].innerHTML = '<input type="button" onclick="deleteEnrolled(this)" value="Delete" id="button-1"/>'
    resetForm()
}

function resetForm() {
    document.getElementById("sid").value = "";
    document.getElementById("cid").value = "";
}

function getStudents(data) {
    const studentIds = []
    const students = JSON.parse(data)
    students.forEach(stud => {
        studentIds.push(stud.StdID)
    });
    var select = document.getElementById("sid")
    //for loops to iterate through the student id and display it
    for (var i = 0; i < studentIds.length; i++){
        var sidd = studentIds[i]
        var option = document.createElement("option")
        option.innerHTML = sidd
        option.value = sidd
        select.appendChild(option)
    }
}

function showCourses(data) {
    const courseIds = []
    const courses = JSON.parse(data)
    courses.forEach(element => {
        courseIds.push(element.id)
    });
    var select = document.getElementById("cid")
    for (var i = 0; i < courseIds.length; i++){
        var cidd = courseIds[i]
        var option = document.createElement("option")
        option.innerHTML = cidd
        option.value = cidd
        select.appendChild(option)
    }
}


function deleteEnrolled(r) {
    selectedRow = r.parentElement.parentElement
    sid = selectedRow.cells[0].innerHTML
    cid = selectedRow.cells[1].innerHTML
    if (confirm(`Are you sure you want to delete the enrollment of student ${sid} from the course ${cid}` )){
        fetch("/deleteEnrolled/"+sid+"/"+cid,{
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