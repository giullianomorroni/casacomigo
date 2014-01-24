/**
* Created by giulliano on 10/18/13.
*/
$(document).ready(function() {
    //Validate (http://jqueryvalidation.org/documentation/)
    $("#assinatura_passo_1").validate({
        errorClass: "erroAssinaturaPasso1",
        rules: {
            casamento: {required:true},
            senha: {required:true, minlength: 6},
            noivo: {required:true},
            noiva: {required:true},
            emailNoiva: {email: true},
            emailNoivo: {email: true}
        },
        messages: {
            noivo: "Qual o nome do noivo ?",
            noiva: "Qual o nome do noiva ?",
            casamento: "Qual a data do casamento ?",
            senha: {
                required: "Digite uma senha ",
				minlength: "Sua senha deve ter no mínimo 6 caracteres"
            },
            emailNoiva: "Email inválido",
            emailNoivo: "Email inválido"
        }
    });

    //http://igorescobar.github.io/jQuery-Mask-Plugin/
    $("#casamento").mask("00/00/0000");
    $("#telefoneNoivo").mask("(00) 0000-00009");
    $("#telefoneNoiva").mask("(00) 0000-00009");

});
