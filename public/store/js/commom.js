var HOST = location.host;
var PROTOCOL = window.location.protocol;

function addToKart() {
	var product_code = $("#product_code").val()
	var quantity = $("#quantity").val()
	
	if (!product_code){
		return;
	}
	
	if (!quantity) {
		quantity=1;
	}

	$.ajax({
		  url: (PROTOCOL+"//"+HOST)+"/carrinho/adicionar/produto/"+product_code+"/quantidade/"+quantity,
		  dataType: "json"
	});
}