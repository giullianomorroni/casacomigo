$( document ).ready(function() {
 
	product_code = $("#product_code").val()
	$.ajax({
		
		  url: "http://localhost:9000/produto/"+product_code,
		  success: function(data) { mountProduct(data); },
		  dataType: "json"
	});
});

function mountProduct(data) {
	$("#title").html(data.Titulo);
	$("#price").html("R$ " + data.Preco);
	$("#descricao_curta").html(data.Descricao.substring(0, 200) + "...");
	$("#descricao_completa").html(data.Descricao);

	for (var x in data.Informacoes){
		info = data.Informacoes[x];
		for (var key in info){
			$("#informacoes").append("<tr class='techSpecRow'><td class='techSpecTD1'> <b>"+ key +" </b> </td><td class='techSpecTD2'>"+info[key]+"</td></tr>");
		}
	}

	for(var x=1; x < 3; x++){
		var urlPhoto = "<a href='/public/store/img/"+data.Codigo+"/"+x+".jpg' alt=''> <img style='width:29%' src='/public/store/img/"+data.Codigo+"/"+x+".jpg' alt=''/></a>";
		$("#photos").append(urlPhoto);
	}

	var mainPhoto = "<a href='/public/store/img/"+data.Codigo+"/thumb.jpg'> <img src='/public/store/img/"+data.Codigo+"/thumb.jpg' style='width:100%'/> </a>";
	$("#mainPhoto").append(mainPhoto);
}
