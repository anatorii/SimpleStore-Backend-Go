alter table clients add
  CONSTRAINT fk_address 
  FOREIGN KEY (address_id)
  REFERENCES addresses(id) ON DELETE SET NULL;

alter table suppliers add
  CONSTRAINT fk_address 
  FOREIGN KEY (address_id)
  REFERENCES addresses(id) ON DELETE SET NULL;

alter table products add
  CONSTRAINT fk_supplier 
  FOREIGN KEY (supplier_id)
  REFERENCES suppliers(id) ON DELETE SET NULL;

alter table products add
  CONSTRAINT fk_image
  FOREIGN KEY (image_id)
  REFERENCES images(id) ON DELETE SET NULL;
