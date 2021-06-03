<?php

class Person
{
  public function __construct(
    private string $firstName,
    private string $lastName,
    private int $birthYear
  ) {
  }

  public function uniqueId(): string
  {
    return $this->firstName[0] . $this->lastName . $this->birthYear;
  }
}

$gary = new Person("Gary", "Lime", 1982);

$id = $gary->uniqueId();

print($id);
