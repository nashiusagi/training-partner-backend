set -ue

sqlite3 resources/training_partner.db < resources/sql/01_create_tables.sql
sqlite3 resources/training_partner.db < resources/sql/02_seeding.sql
