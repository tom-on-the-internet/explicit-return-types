class Person {
  constructor(
    private firstName: string,
    private lastName: string,
    private birthYear: number
  ) {}

  uniqueId(): string {
    return this.firstName[0] + this.lastName + this.birthYear;
  }
}

const gary = new Person("Gary", "Lime", 1982);

const id = gary.uniqueId();

console.log(id);
