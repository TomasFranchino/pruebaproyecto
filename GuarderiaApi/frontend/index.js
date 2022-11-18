const user = document.getElementById ('username')
const password = document.getElementById('password')
const button = document.getElementById('button')
button.addEventListener('click', (e) => { e.preventDefault()
    const data = {
        user: user.value, 
        password: password.value 
}
const API_URL =  'http://localhost:8080/api/v1/socio/'+data.user+'/'+data.password+'';
login(API_URL);
})

const login = async(API_URL) => {

    try{
        const respuesta = await fetch(`${API_URL}`);

        const datos = await respuesta.json();
        if (datos == true){
          window.location= "/viewes/user/calculadora.html";
        }else{
            document.getElementById("estado-de-inicio").innerHTML = "Usuario o contraseña incorrecta.";
            
        
        }
}
catch(error){
}
}

