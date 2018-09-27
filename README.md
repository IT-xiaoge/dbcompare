### A Simple DbCompare for mysql

```json
{
  "DbTables": [
    [ "table_exist_in_db1_dn" ], // These  tables exist in Db1Dn and not exist in Db2Dn
    [  "table_exist_in_db2_dn" ] // These  tables exist in Db2Dn and not exist in Db1Dn
  ],
  "TablesDiffResult": {
    "usr_login": {
      "Fields": [ 
            ["field_exist_in_db1_dn"], // These  fields exist in Db1Dn and not exist in Db2Dn
            ["field_exist_in_db2_dn"]  // These  fields exist in Db2Dn and not exist in Db1Dn
      ],
      "FieldDiffResult": {
        "create_time": {
          "Type": [ 
          "int(10) unsigned",   // Field type in Db1Dn
          "int(10)"    // Field type in Db2Dn
          ],
          "Null": [ 
          "NO",  // Field allow null in Db1Dn
          "NO"   // Field allow null in Db1Dn
          ],
          "Key": [ 
           "", // Field key in Db1Dn
           "" // Field key in Db1Dn
           ],
          "Default": [
           "0", // Field default value in Db1Dn
           "0" // Field default value in Db1Dn
            ],
          "Extra": [ 
          "", // Field extra in Db1Dn
          ""// Field extra in Db1Dn
           ]
        },
        "id": {
          "Type": [ "bigint(20) unsigned", "bigint(20) unsigned" ],
          "Null": [  "NO", "NO" ],
          "Key": [ "PRI", "PRI" ],
          "Default": [ "", "" ],
          "Extra": [ "", "" ]
        }
      }
    }
  }
}
```