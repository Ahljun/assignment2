<html>
    <header>
	    <title>iLift</title>
	</header>
    <body>
    
    <form action="/sign" method="post">
        <div>Height(cm) <textarea name="height" rows="1" cols="3"></textarea></div><br>
	    <div>Weight(kg) <textarea name="weight" rows="1" cols="3"></textarea></div><br>
	  
        <div><input type="submit" value="Submit" name="makeResults"></div>
    </form>
	BMI = <?php echo $BMI; ?>
    </body>
</html>

<?php
    $height = ($POST['height']);
    $weight = ($POST['weight']);
    $BMI = ($height + $weight);
?>



