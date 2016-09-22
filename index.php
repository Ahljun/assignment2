<html>
<head>
    <title>iLift</title>
</head>
<body>

<form action="" method="post">

    <label>Weight(kg):</label>
    <input type="number" name="weight" min="0"/><br><br>

    <label>Height(m):</label>
    <input type="number" name="height" step="0.01" min="0"/><br><br>

    <input  type="submit" name="btn_submit" value="Submit">
</form>

<?php
    if(isset($_POST['btn_submit']))
    {
        $weight = $_POST['weight'];
        $height = $_POST['height'];
        
        $BMI =  ($weight/($height*$height));
        
        echo "Your BMI is: ".$BMI;
    }
?>


</body>
</html>