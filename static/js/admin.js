var form = document.querySelector('form');

var imgs = [];
var counter =0;

var ch = function(e) {
    e.preventDefault();
    var x = new XMLHttpRequest()

    x.onreadystatechange = function() {
        if(x.readyState == 4) {
            if(x.responseText !=""){
                imgs.push("data:image/png;base64," + x.responseText + "\ ");
                var test = document.getElementById('imgs');
                test.innerHTML += '<img id= '+ counter +' onclick="del(event)" src=\ '+ imgs[counter] + '/>';
                counter++;
            }
            console.log(imgs.length);
        }
    }
    x.open("POST", "temp-post");
    x.send(new FormData(form));
};

var del = function(e) {
    var x = new XMLHttpRequest()
    var index = e.target.id
    imgs.splice(index, 1);
    counter--;
    console.log(imgs.length);

    var parent = document.getElementById('imgs');
    var img = document.getElementById(index);
    parent.removeChild(img);
}

var upload = function(e){
    e.preventDefault();

    var x = new XMLHttpRequest();
    var jsn = JSON.stringify(imgs);
    x.open("POST", "upload");
    x.send(jsn);
}