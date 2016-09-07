<?php 

fscanf(STDIN, "%d\n", $t);
for($i = 0; $i < $t; $i++) {
		fscanf(STDIN, "%s\n", $numString); 
		$flag = true; 
		$countZero = 0; 
		$countOne = 0;
		$count = strlen($numString);

		for ($j=0; $j < $count; $j++) { 		
			if($numString[$j] == 0) { 
				$countZero ++ ;
			} else { 
				$countOne ++; 
			} 
		}

		if ($count != 1 &&  ($countZero == 0 || $countOne == 0 || ($countZero > 1 && $countOne > 1))) {		
					$flag = false;
		}  

		if ($flag) { 
			printf("%s\n", "Yes"); 
		} else { 
			printf("%s\n", "No");
		} 
} 