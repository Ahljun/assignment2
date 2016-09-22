<html>
  
  <head>
    <title>iLift</title>
  </head>

  <?php
  
      $handle = fopen ('gs://s3485849-ass2-storage/bmi.txt','w');
      fwrite ($handle, "2");
  
      if(isset($_POST['btn_submit']))
      
      {
          $weight = $_POST['weight'];
          $height = $_POST['height'];
          
          $BMI =  ($weight/($height*$height));
      }
      
      fwrite ($handle, ", " . $BMI);
      break;  
      
  ?>
  
  <body>
  
    <form action="" method="post">
    
        <label>Weight(kg):</label>
        <input type="number" name="weight" rows="1" cols="3"/><br><br>
    
        <label>Height(cm):</label>
        <input type="number" name="height" step="0.01" rows="1" cols="3"/><br><br>
    
        <input  type="submit" name="btn_submit" value="Submit">
        
    </form>
    
     <?php echo "The total value is: ".$BMI; ?>
  
  </body>
  
</html>