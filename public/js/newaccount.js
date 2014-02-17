$( document ).ready(function() { 
		$("#assinatura_passo_1").validate({
			rules: {
				nomeNoivo : { required: true },
				nomeNoiva : { required: true },
				//telefoneNoivo : { required: true },
				//telefoneNoiva : { required: true},
				emailNoivo : { required: true, email: true },
				emailNoiva : { required: true, email: true },
				apelido : { required: true },
				senha : { required: true, minlength: 6 }
			},
			messages: { 
				nomeNoivo : "Este campo é obrigatório",
				nomeNoiva : "Este campo é obrigatório",
				//telefoneNoivo : "Este campo é obrigatório",
				//telefoneNoiva : "Este campo é obrigatório",
				emailNoivo : { required: "Este campo é obrigatório", email: "O formato de email inválido" },
				emailNoiva : { required: "Este campo é obrigatório", email: "O formato de email inválido" },
				apelido : "Este campo é obrigatório",
				senha : { required: "Este campo é obrigatório", minlength: "A senha deve ter pelo menos 6 caracteres"}
			}
		});
});