var form = document.querySelector('form');

var imgs = [];
var sendImgs = [];
var counter =0;

var ch = function(e) {
    e.preventDefault();
    var x = new XMLHttpRequest()

    x.onreadystatechange = function() {
        if(x.readyState == 4) {
            if(x.responseText !=""){
                imgs.push("data:image/png;base64," + x.responseText + "\ ");
                sendImgs.push(x.responseText);
                var img = document.getElementById('imgs');
                img.innerHTML += '<img id= '+ counter +' onclick="del(event)" src=\ '+ imgs[counter] + '/>';
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
    sendImgs.splice(index, 1);
    counter--;
    console.log(imgs.length);

    var parent = document.getElementById('imgs');
    var img = document.getElementById(index);
    parent.removeChild(img);
}

var upload = function(e){
    e.preventDefault();

    var x = new XMLHttpRequest();
    var jsn = JSON.stringify(sendImgs);
    x.onreadystatechange = function() {
        if(x.readyState ==4 ){
            if(x.responseText == ""){
                window.location.replace('/admin/portfolio');
            }
        }
    }
    x.open("POST", "upload");
    x.setRequestHeader('Content-type', 'application/json');
    x.send(jsn);
}

var deletePic = function(e) {
    e.preventDefault();
    var elemet = e.currentTarget.children[0].src;
    elemet = elemet.split("/portfolio")
    elemet =elemet[1].split("?")
    
    var cont = confirm("Are you sure you want to delete this picture? ");
    if(!cont) {
        return;
    }
    
    var x = new XMLHttpRequest();
    x.onreadystatechange = function() {
        if(x.readyState ==4 ){
            location.reload(true);
            
        }
    }
    x.open("POST", "portfolio/delete");
    x.send(elemet[0]);
}

var contactRow = function(e) {
    console.log(e);
    window.location.href = '/admin/contact/' + e;

}

