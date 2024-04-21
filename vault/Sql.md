#todo #sql

#### Common things
##### ACID
> ACID stands for Atomicity, Consistency, Isolation, and Durability. These are the four key properties that ensure the reliability of database transactions:
> 1. Atomicity: This property ensures that a transaction is treated as a single unit of work, meaning that either all of its operations are successfully completed and committed, or none of them are. There are no partial transactions.
> 2. Consistency: This property ensures that the database remains in a consistent state before and after the transaction. It guarantees that all constraints, rules, and relationships are maintained, preserving the integrity of the data.
> 3. Isolation: Isolation ensures that the concurrent execution of transactions results in a state that would be obtained if the transactions were executed sequentially. It prevents interference between transactions, ensuring that each transaction operates independently of and transparently to other transactions.
> 4. Durability: Durability ensures that once a transaction is committed, its changes are permanently stored in the database and will survive system failures. Even in the event of a crash or power failure, the changes made by committed transactions will not be lost.
>
> Together, these properties ensure that database transactions are reliable, maintain data integrity, and provide a high level of data consistency and availability.

##### CRUD 
> Crud stands for Create, Read, Update, and Delete. It represents the four basic operations  that can be performed on data in most relational database management systems (RDBMS) or in any persistent storage system:
> 1. **Create**: This operation involves creating new records or entries in a database. It typically corresponds to the insertion of new data into a database table.
> 2. **Read**: This operation involves retrieving or querying existing data from a database. It allows users to access and view stored data without modifying it.
> 3. **Update**: This operation involves modifying existing data in a database. It allows users to change the values of specific attributes of a record or entry.
> 4. **Delete**: This operation involves removing existing data from a database. It allows users to permanently remove records or entries from a database table.
>
>CRUD operations are fundamental to data manipulation and are commonly used in applications to interact with databases. They form the basis for most database-driven applications' functionalities, allowing users to create, retrieve, update, and delete data as needed.

##### N+1 problem
> For example, you have ***cars*** table and related ***wheels*** table, so each car has many wheels ([[TODO|one to many]] relation). Then you decide to get cars: 
```sql
SELECT * FROM cars;
```
> and later (maybe for because of another trigger) you join them their wheels:
```sql
SELECT * FROM wheels WHERE wheels.id = $car.id;
``` 
> then you do 1 (cars select) + N (select wheels * cars number) queries, which can hurt performance

##### CAP theorem
> In [theoretical computer science](https://en.wikipedia.org/wiki/Theoretical_computer_science "Theoretical computer science"), the **CAP theorem**, also named **Brewer's theorem** after computer scientist [Eric Brewer](https://en.wikipedia.org/wiki/Eric_Brewer_(scientist) "Eric Brewer (scientist)"), states that any [distributed data store](https://en.wikipedia.org/wiki/Distributed_data_store "Distributed data store") can provide only [two of the following three](https://en.wikipedia.org/wiki/Trilemma "Trilemma") guarantees:


> 1. [Consistency](https://en.wikipedia.org/wiki/Consistency_model "Consistency model")
Every read receives the most recent write or an error.

> 2. [Availability](https://en.wikipedia.org/wiki/Availability "Availability")
Every request receives a (non-error) response, without the guarantee that it contains the most recent write.

> 3. [Partition tolerance](https://en.wikipedia.org/wiki/Network_partitioning "Network partitioning")
The system continues to operate despite an arbitrary number of messages being dropped (or delayed) by the network between nodes.

>When a [network partition](https://en.wikipedia.org/wiki/Network_partition "Network partition") failure happens, it must be decided whether to do one of the following:
> - cancel the operation and thus decrease the availability but ensure consistency
> - proceed with the operation and thus provide availability but risk inconsistency.