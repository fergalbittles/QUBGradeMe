<?php
if (!function_exists("getSortedModules")) {
  function getSortedModules($modules, $marks) {
    // Ensure both params are actually arrays
    if(!is_array($modules) || !is_array($marks)) {
      return null;
    }
  
    // Ensure neither arrays are empty
    if (count($modules) == 0 || count($marks) == 0) {
      return null;
    }
  
    $module_marks = array();
    for ($i = 0; $i < count($modules); $i++) {
      // Ensure the module/mark at this index is not empty or null (marks are allowed to be 0)
      if (empty($modules[$i]) || (empty($marks[$i]) && $marks[$i] != 0) || empty(trim($modules[$i]))) {
        return null;
      }
  
      // Ensure the mark at this index is actually a number
      if (!is_int($marks[$i])) {
        return null;
      } 
    
      // Ensure the mark is within suitable range
      if ($marks[$i] < 0 || $marks[$i] > 100) {
        return null;
      }
  
      $module_marks_array = array("module"=>$modules[$i], "marks"=>$marks[$i]);
      array_push($module_marks,$module_marks_array);
    }
  
    usort($module_marks, function($a, $b) {
      return $b['marks'] <=> $a['marks'];
    });
  
    return $module_marks;
  }
}

