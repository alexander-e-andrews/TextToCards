var input
var preview
var fileInput
var csvInput


window.onload = function () {
    input = document.querySelector('input');
    preview = document.getElementById('preview');
    fileInput = document.getElementById("background")
    csvInput = document.getElementById("csvSource")
    console.log(fileInput)
    fileInput.addEventListener('change', updateImageDisplay);
}

function updateImageDisplay() {
    console.log(fileInput)
    file = fileInput.files[0]
    console.log(file)
    preview.src = URL.createObjectURL(file);
}

const fileTypes = [
    'image/jpeg',
    'image/pjpeg',
    'image/png'
];
files[""0""]
function validFileType(file) {
    return fileTypes.includes(file.type);
}

function returnFileSize(number) {
    if (number < 1024) {
        return number + 'bytes';
    } else if (number >= 1024 && number < 1048576) {
        return (number / 1024).toFixed(1) + 'KB';
    } else if (number >= 1048576) {
        return (number / 1048576).toFixed(1) + 'MB';
    }
}

function PreviewCard() {
    console.log("Preview Card Called")
    var xhr = new XMLHttpRequest();

    //Ontrutnr function
    xhr.onload = function () {
        console.log("Got a response")
    }

    console.log("Hello")
    xhr.open("POST", "/previewCard")
    var formData = new FormData();

    background = fileInput.files[0]
    csv = csvInput.files[0]

    formData.append("background", background)
    formData.append("csv", csv)
    formData.append("child", "noProblem")

    //Send the proper header information along with the request
    //xhr.setRequestHeader("Content-Type", " multipart/form-data")
    xhr.responseType = "json";
    xhr.send(formData);
}

function GenerateCards() {

}