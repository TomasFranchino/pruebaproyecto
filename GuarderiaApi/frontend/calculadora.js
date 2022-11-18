const txtoperador1 = document.getElementById("operador1")
const txtoperador= document.getElementById("operador")
const txtoperdor2= document.getElementById("operador2")
const btncalcular = document.getElementById("calcular")
const presultado= document.getElementById("resultado")

btncalcular.addEventListener('click',Calcular)

function Calcular(){
    const op =txtoperador.value
    const op1=parseFloat(txtoperador1.value)
    const op2=parseFloat(txtoperdor2.value)

    if((op== "+" || op== "-" ||op== "*" ||op== "/")&&!isNaN(op1)&&!isNaN(op2)){
        presultado.innerText="Calculo posible"
        let resultado
        switch(op){
            
            case "+":
                resultado=op1+op2
                break
            case "-":
                resultado=op1-op2
                break
            case "*":
                resultado=op1*op2
                break
            case "/":
                resultado=op1/op2
                break

            

        }
        presultado.innerText= "="+resultado
    }else{
        presultado.innerText="Calculo imposible"
    }


}