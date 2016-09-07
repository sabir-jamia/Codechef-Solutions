<?php 

fscanf(STDIN, "%d\n", $t);
for($i = 0; $i < $t; $i++) {
		fscanf(STDIN, "%s\n", $str); 
		$flag = true;
		$count = strlen($numString);

		if($count % 2 != 0 ) {
			$len = ($count - 1)/ 2;
		} else {
			$len = $count / 2;
		}

		for ($j=0; $j < $len; $j++) { 		
			if($str[$j] == 0 && $str[$count - $j - 1]) {
				$flag = false;
				break;
			}
		}

		if ($flag) { 
			for ($j=0; $j < $count; $j++) { 		
				if($str[$j] == 0 && $str[$count - $j - 1]) {
					$flag = false;
				break;
			}
		}
			printf("%s\n", ""); 
		} else { 
			printf("%s\n", "-1");
		} 
} 