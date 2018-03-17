var x = document.querySelector('#t')
x.onclick = open;
function open(){
    document.getElementById('nav').classList.toggle('active');
    document.getElementById('header').classList.toggle('active')
    document.getElementById('t').classList.toggle('open');
    document.body.classList.toggle('open');
}

function hideMsg() {
    document.getElementById('login-msg').classList.add('hidden');
}
