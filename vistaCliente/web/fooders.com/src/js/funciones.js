(function($) {
    
  
    
    // Activate scrollspy to add active class to navbar items on scroll
    $('body').scrollspy({
      target: '#mainNav',
      offset: 54
    });

    //$('.menus-contenido').html("Hola soy miguel xd");
    var info;
    $.ajax({
        url:'/js/ajax/api.php',
        type:"GET",
        data:{
            nameTable:"restaurantes",
        },
        success:function(data){
            console.log(data);
            info = JSON.parse(data);
            tam= info.length;
            i=0;
            while(i<tam){
                console.log(info[i]['nombre']);
                i++;
            }
            
        }
    });
  
  })(jQuery); // End of use strict


$(document).ready(function(){
    //url = 'https://c5u2jk30nd.execute-api.us-west-2.amazonaws.com/nuevo/platos?callback=?';
    //datos = {};
    //$("#buscaPorRestaurante").on('show.bs.modal', function(evnt){
      //  consola.log("Modal activado");
    //});

    $.ajax({
        url:'/js/ajax/api.php',
        type:"GET",
        data:{
            nameTable:"restaurantes"
        },
        //async:false,
        success:function(data){
            var array = JSON.parse(data);
            tam= array.length;
            content_tabla="<table class='table table-responsive table-bordered table-striped border'>";
            content_tabla+="<thead class='thead-dark'><tr>";
            content_tabla+="<th class='text-center align-items-center align-middle'>N°</th>";
            content_tabla+= "<th class='text-center col-lg align-items-center align-middle'>Nombre Restaurante</th>";
            content_tabla+= "<th class='text-center align-items-center align-middle'>latitud</th>";
            content_tabla+= "<th class='text-center align-items-center align-middle'>longitud</th>";
            content_tabla+="</tr>";
            content_tabla+="</thead>";
            
            i=0;
            while(i<tam){
                content_tabla+= "<tr>";
                content_tabla+="<td class='text-center align-items-center align-middle'>"+parseInt(i+1)+"</td>";
                content_tabla+="<td class='text-center align-items-center align-middle'>"+array[i]['nombre']+"</td>";
                content_tabla+="<td class='text-center align-items-center align-middle'>"+array[i]['latitud']+"</td>";
                content_tabla+="<td class='text-center align-items-center align-middle'>"+array[i]['longitud']+"</td>";
                content_tabla+= "</tr>";
                i++;
            }
                
            content_tabla+="</table>";

            $("#tablaCoordenadas").html(content_tabla);
        }
    });

    $("#buscaPorRestaurante").on('show.bs.modal',function(e){

        //console.log("Estoy aqui");
        $.ajax({
            url:'/js/ajax/escogerRestaurante.php',
            type:"GET",
            data:{
                //nameTable:"restaurantes",
            },
            success:function(data){
                console.log(data);
                var info = JSON.parse(data);
                tam= info.length;
                content_html="<option value=''>Seleccion un Restuarante</option>";
                
                i=0;
                while(i<tam){
                    //console.log(info[i]['nombre']);
                    content_html+="<option value='"+info[i]['id']+"'>"+info[i]['nombre']+"</option>";
                    i++;
                    $("#cb_restaurante").html(content_html);
                    $("#cb_resultRestaurante").html("xxxxxxxxxxxxxxxxxx");
                }
                //console.log(content_html);
                
            }
        });
    
        $("#cb_restaurante").change(function(){
            $("#cb_resultRestaurante").remove();
            
            $("#cb_restaurante option:selected").each(function(){
                id_restaurante=$(this).val();
                console.log("se escogio el restaurante id ->"+id_restaurante);
                
            });
        });

    });

    
});


function listaPlatos(){
    $(document).ready(function(){

        //$.ajax();
        $(".mostrar_buscaPorrestaurante");
    });
}