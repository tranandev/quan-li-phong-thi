//SPDX-License-Identifier: Apache-2.0

var phongthi = require('./controller.js');

module.exports = function(app){

  app.get('/tim_ID/:id', function(req, res){
    phongthi.tim_ID(req, res);
  });
  app.get('/tao_ID/:phongthi', function(req, res){
    phongthi.tao_ID(req, res);
  });
  app.get('/tim_tat_ca_ID', function(req, res){
    phongthi.tim_tat_ca_ID(req, res);
  });
  app.get('/change_diemthi/:diemthi', function(req, res){
    phongthi.change_diemthi(req, res);
  });
}
