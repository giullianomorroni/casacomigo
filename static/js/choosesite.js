$( document ).ready(function() { 
		$("#assinatura_passo_2").validate({
			ignore: "",
			rules: {
				site_id : { required: true }
			},
			messages: { 
				site_id : "Escolha um modelo antes de prosseguir"
			}
		});
});